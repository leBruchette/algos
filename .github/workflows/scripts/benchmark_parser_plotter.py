#!/usr/bin/env python3
"""
Benchmark results parser and plotter for Go benchmarks.
File: scripts/plot_benchmarks.py
"""

import re
import os
import json
import datetime
import matplotlib.pyplot as plt
import matplotlib.dates as mdates
import pandas as pd
import numpy as np
import seaborn as sns
from pathlib import Path

# Set up matplotlib for headless operation
plt.switch_backend('Agg')
sns.set_style("whitegrid")

class BenchmarkParser:
    def __init__(self, benchmark_file="benchmark-results/benchmark_output.txt"):
        self.benchmark_file = benchmark_file
        self.results = []
        
    def parse_benchmark_output(self):
        """Parse Go benchmark output into structured data."""
        benchmark_pattern = r'Benchmark(\w+)-(\d+)\s+(\d+)\s+([\d.]+)\s+ns/op(?:\s+([\d.]+)\s+B/op)?(?:\s+([\d.]+)\s+allocs/op)?'
        
        with open(self.benchmark_file, 'r') as f:
            content = f.read()
            
        matches = re.findall(benchmark_pattern, content)
        
        for match in matches:
            benchmark_name, cpus, iterations, ns_per_op, bytes_per_op, allocs_per_op = match
            
            result = {
                'name': benchmark_name,
                'cpus': int(cpus),
                'iterations': int(iterations),
                'ns_per_op': float(ns_per_op),
                'bytes_per_op': float(bytes_per_op) if bytes_per_op else 0,
                'allocs_per_op': float(allocs_per_op) if allocs_per_op else 0,
                'timestamp': datetime.datetime.now().isoformat()
            }
            
            self.results.append(result)
            
        return self.results
    
    def group_by_benchmark(self):
        """Group results by benchmark name for averaging multiple runs."""
        grouped = {}
        
        for result in self.results:
            name = result['name']
            if name not in grouped:
                grouped[name] = []
            grouped[name].append(result)
            
        # Average multiple runs
        averaged = {}
        for name, runs in grouped.items():
            if len(runs) > 1:
                avg_result = {
                    'name': name,
                    'cpus': runs[0]['cpus'],
                    'iterations': sum(r['iterations'] for r in runs) // len(runs),
                    'ns_per_op': sum(r['ns_per_op'] for r in runs) / len(runs),
                    'bytes_per_op': sum(r['bytes_per_op'] for r in runs) / len(runs),
                    'allocs_per_op': sum(r['allocs_per_op'] for r in runs) / len(runs),
                    'timestamp': runs[0]['timestamp'],
                    'run_count': len(runs)
                }
                averaged[name] = avg_result
            else:
                averaged[name] = runs[0]
                
        return averaged

class BenchmarkPlotter:
    def __init__(self, results):
        self.results = results
        self.output_dir = Path("benchmark-results/plots")
        self.output_dir.mkdir(parents=True, exist_ok=True)
        
    def plot_performance_comparison(self):
        """Create a bar chart comparing benchmark performance."""
        if not self.results:
            return
            
        names = list(self.results.keys())
        ns_per_op = [self.results[name]['ns_per_op'] for name in names]
        
        fig, ax = plt.subplots(figsize=(12, 8))
        
        # Create bars with different colors
        colors = plt.cm.viridis(np.linspace(0, 1, len(names)))
        bars = ax.bar(names, ns_per_op, color=colors)
        
        # Add value labels on bars
        for bar, value in zip(bars, ns_per_op):
            height = bar.get_height()
            ax.text(bar.get_x() + bar.get_width()/2., height,
                   f'{value:.0f}', ha='center', va='bottom')
        
        ax.set_ylabel('Nanoseconds per Operation')
        ax.set_title('Go Benchmark Performance Comparison')
        ax.set_yscale('log')  # Use log scale for better visualization
        
        # Rotate x-axis labels for better readability
        plt.xticks(rotation=45, ha='right')
        plt.tight_layout()
        
        plt.savefig(self.output_dir / 'performance_comparison.png', dpi=300, bbox_inches='tight')
        plt.close()
        
    def plot_memory_usage(self):
        """Create a chart showing memory usage per operation."""
        memory_data = {name: data for name, data in self.results.items() 
                      if data['bytes_per_op'] > 0}
        
        if not memory_data:
            return
            
        names = list(memory_data.keys())
        bytes_per_op = [memory_data[name]['bytes_per_op'] for name in names]
        allocs_per_op = [memory_data[name]['allocs_per_op'] for name in names]
        
        fig, (ax1, ax2) = plt.subplots(1, 2, figsize=(15, 6))
        
        # Bytes per operation
        ax1.bar(names, bytes_per_op, color='skyblue')
        ax1.set_ylabel('Bytes per Operation')
        ax1.set_title('Memory Usage per Operation')
        ax1.tick_params(axis='x', rotation=45)
        
        # Allocations per operation
        ax2.bar(names, allocs_per_op, color='lightcoral')
        ax2.set_ylabel('Allocations per Operation')
        ax2.set_title('Memory Allocations per Operation')
        ax2.tick_params(axis='x', rotation=45)
        
        plt.tight_layout()
        plt.savefig(self.output_dir / 'memory_usage.png', dpi=300, bbox_inches='tight')
        plt.close()
        
    def plot_scalability_analysis(self):
        """Analyze benchmark scalability by looking for size patterns in names."""
        # Group benchmarks by base name (removing size suffixes)
        scalability_groups = {}
        
        for name, data in self.results.items():
            # Try to extract size from benchmark name (e.g., BenchmarkSort_1000)
            base_name = re.sub(r'_\d+$', '', name)
            size_match = re.search(r'_(\d+)$', name)
            
            if size_match:
                size = int(size_match.group(1))
                if base_name not in scalability_groups:
                    scalability_groups[base_name] = []
                scalability_groups[base_name].append((size, data['ns_per_op']))
        
        if not scalability_groups:
            return
            
        fig, ax = plt.subplots(figsize=(12, 8))
        
        colors = plt.cm.Set1(np.linspace(0, 1, len(scalability_groups)))
        
        for i, (base_name, size_data) in enumerate(scalability_groups.items()):
            size_data.sort()  # Sort by size
            sizes, times = zip(*size_data)
            
            ax.plot(sizes, times, marker='o', label=base_name, color=colors[i])
            
        ax.set_xlabel('Input Size')
        ax.set_ylabel('Nanoseconds per Operation')
        ax.set_title('Benchmark Scalability Analysis')
        ax.set_xscale('log')
        ax.set_yscale('log')
        ax.legend()
        ax.grid(True, alpha=0.3)
        
        plt.tight_layout()
        plt.savefig(self.output_dir / 'scalability_analysis.png', dpi=300, bbox_inches='tight')
        plt.close()
        
    def create_summary_report(self):
        """Create a markdown summary report."""
        summary_lines = ["# Benchmark Results Summary\n"]
        summary_lines.append(f"**Generated:** {datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S UTC')}\n")
        summary_lines.append(f"**Total Benchmarks:** {len(self.results)}\n")
        
        # Performance table
        summary_lines.append("## Performance Results\n")
        summary_lines.append("| Benchmark | ns/op | B/op | allocs/op | Iterations |")
        summary_lines.append("|-----------|-------|------|-----------|------------|")
        
        for name, data in sorted(self.results.items(), key=lambda x: x[1]['ns_per_op']):
            summary_lines.append(
                f"| {name} | {data['ns_per_op']:.0f} | {data['bytes_per_op']:.0f} | "
                f"{data['allocs_per_op']:.2f} | {data['iterations']:,} |"
            )
        
        # Best/worst performers
        if self.results:
            fastest = min(self.results.items(), key=lambda x: x[1]['ns_per_op'])
            slowest = max(self.results.items(), key=lambda x: x[1]['ns_per_op'])
            
            summary_lines.append(f"\n## Key Insights\n")
            summary_lines.append(f"- **Fastest:** {fastest[0]} ({fastest[1]['ns_per_op']:.0f} ns/op)")
            summary_lines.append(f"- **Slowest:** {slowest[0]} ({slowest[1]['ns_per_op']:.0f} ns/op)")
            
            if fastest[1]['ns_per_op'] > 0:
                speedup = slowest[1]['ns_per_op'] / fastest[1]['ns_per_op']
                summary_lines.append(f"- **Speed Difference:** {speedup:.1f}x")
        
        summary_lines.append("\n## Generated Charts\n")
        summary_lines.append("- [Performance Comparison](performance_comparison.png)")
        summary_lines.append("- [Memory Usage](memory_usage.png)")
        summary_lines.append("- [Scalability Analysis](scalability_analysis.png)")
        
        with open("benchmark-results/summary.md", "w") as f:
            f.write("\n".join(summary_lines))

def main():
    print("Parsing benchmark results...")
    parser = BenchmarkParser()
    
    if not os.path.exists(parser.benchmark_file):
        print(f"Benchmark file {parser.benchmark_file} not found!")
        return
        
    results = parser.parse_benchmark_output()
    
    if not results:
        print("No benchmark results found!")
        return
        
    print(f"Found {len(results)} benchmark results")
    
    # Group and average multiple runs
    averaged_results = parser.group_by_benchmark()
    
    # Create plots
    print("Generating plots...")
    plotter = BenchmarkPlotter(averaged_results)
    
    plotter.plot_performance_comparison()
    plotter.plot_memory_usage()
    plotter.plot_scalability_analysis()
    plotter.create_summary_report()
    
    # Save raw data as JSON for historical tracking
    with open("benchmark-results/benchmark_data.json", "w") as f:
        json.dump(averaged_results, f, indent=2)
    
    print("✅ Benchmark analysis complete!")
    print(f"📊 Plots saved to: benchmark-results/plots/")
    print(f"📝 Summary saved to: benchmark-results/summary.md")

if __name__ == "__main__":
    main()