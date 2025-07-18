#!/usr/bin/env python3
"""
Historical benchmark tracking and trend analysis.
File: scripts/update_benchmark_history.py
"""

import json
import os
import datetime
import matplotlib.pyplot as plt
import matplotlib.dates as mdates
import pandas as pd
from pathlib import Path

class BenchmarkHistoryTracker:
    def __init__(self, history_dir="benchmark-results/history"):
        self.history_dir = Path(history_dir)
        self.history_dir.mkdir(parents=True, exist_ok=True)
        self.history_file = self.history_dir / "benchmark_history.json"
        
    def load_history(self):
        """Load existing benchmark history."""
        if self.history_file.exists():
            with open(self.history_file, 'r') as f:
                return json.load(f)
        return {}
    
    def save_history(self, history):
        """Save benchmark history to file."""
        with open(self.history_file, 'w') as f:
            json.dump(history, f, indent=2)
    
    def update_history(self, current_results):
        """Update history with current benchmark results."""
        history = self.load_history()
        
        # Get current commit info from GitHub environment
        commit_sha = os.getenv('GITHUB_SHA', 'unknown')[:8]
        commit_ref = os.getenv('GITHUB_REF', 'unknown')
        timestamp = datetime.datetime.now().isoformat()
        
        # Create entry for this run
        run_entry = {
            'timestamp': timestamp,
            'commit': commit_sha,
            'ref': commit_ref,
            'results': current_results
        }
        
        # Add to history
        if 'runs' not in history:
            history['runs'] = []
        
        history['runs'].append(run_entry)
        
        # Keep only last 100 runs to avoid huge files
        if len(history['runs']) > 100:
            history['runs'] = history['runs'][-100:]
        
        self.save_history(history)
        return history
    
    def plot_performance_trends(self, history):
        """Plot performance trends over time."""
        if not history.get('runs') or len(history['runs']) < 2:
            print("Not enough historical data for trend analysis")
            return
            
        # Extract benchmark names from the most recent run
        latest_run = history['runs'][-1]
        benchmark_names = list(latest_run['results'].keys())
        
        # Create subplots for each benchmark
        fig, axes = plt.subplots(len(benchmark_names), 1, 
                                figsize=(12, 4 * len(benchmark_names)))
        
        if len(benchmark_names) == 1:
            axes = [axes]
            
        for i, benchmark_name in enumerate(benchmark_names):
            ax = axes[i]
            
            # Extract data for this benchmark across all runs
            timestamps = []
            ns_per_op_values = []
            commits = []
            
            for run in history['runs']:
                if benchmark_name in run['results']:
                    timestamps.append(datetime.datetime.fromisoformat(run['timestamp']))
                    ns_per_op_values.append(run['results'][benchmark_name]['ns_per_op'])
                    commits.append(run['commit'])
            
            if len(timestamps) < 2:
                continue
                
            # Plot the trend
            ax.plot(timestamps, ns_per_op_values, marker='o', linewidth=2, markersize=4)
            ax.set_title(f'{benchmark_name} Performance Trend')
            ax.set_ylabel('ns/op')
            ax.grid(True, alpha=0.3)
            
            # Format x-axis
            ax.xaxis.set_major_formatter(mdates.DateFormatter('%m-%d %H:%M'))
            ax.xaxis.set_major_locator(mdates.DayLocator(interval=1))
            plt.setp(ax.xaxis.get_majorticklabels(), rotation=45)
            
            # Add trend line
            if len(timestamps) >= 3:
                # Convert timestamps to numerical values for regression
                x_numeric = [(ts - timestamps[0]).total_seconds() for ts in timestamps]
                coeffs = np.polyfit(x_numeric, ns_per_op_values, 1)
                trend_line = np.poly1d(coeffs)
                ax.plot(timestamps, trend_line(x_numeric), '--', alpha=0.7, color='red')
                
                # Add trend annotation
                slope = coeffs[0]
                trend_text = "↗️ Slower" if slope > 0 else "↘️ Faster"
                ax.text(0.02, 0.98, trend_text, transform=ax.transAxes, 
                       verticalalignment='top', bbox=dict(boxstyle='round', facecolor='wheat', alpha=0.5))
        
        plt.tight_layout()
        plt.savefig(self.history_dir / 'performance_trends.png', dpi=300, bbox_inches='tight')
        plt.close()
    
    def plot_regression_analysis(self, history):
        """Detect performance regressions between runs."""
        if not history.get('runs') or len(history['runs']) < 2:
            return
            
        current_run = history['runs'][-1]
        previous_run = history['runs'][-2]
        
        regressions = []
        improvements = []
        
        for benchmark_name in current_run['results']:
            if benchmark_name in previous_run['results']:
                current_time = current_run['results'][benchmark_name]['ns_per_op']
                previous_time = previous_run['results'][benchmark_name]['ns_per_op']
                
                change_percent = ((current_time - previous_time) / previous_time) * 100
                
                if abs(change_percent) > 5:  # Only consider significant changes
                    if change_percent > 0:
                        regressions.append((benchmark_name, change_percent, current_time, previous_time))
                    else:
                        improvements.append((benchmark_name, change_percent, current_time, previous_time))
        
        # Create regression report
        if regressions or improvements:
            fig, (ax1, ax2) = plt.subplots(1, 2, figsize=(16, 8))
            
            # Plot regressions
            if regressions:
                names, changes, _, _ = zip(*regressions)
                ax1.barh(names, changes, color='red', alpha=0.7)
                ax1.set_title('Performance Regressions')
                ax1.set_xlabel('Performance Change (%)')
                ax1.grid(True, alpha=0.3)
                
                # Add value labels
                for i, (name, change, current, previous) in enumerate(regressions):
                    ax1.text(change + 0.5, i, f'{change:+.1f}%', 
                            va='center', ha='left')
            else:
                ax1.text(0.5, 0.5, 'No regressions detected', 
                        ha='center', va='center', transform=ax1.transAxes)
                ax1.set_title('Performance Regressions')
            
            # Plot improvements
            if improvements:
                names, changes, _, _ = zip(*improvements)
                ax2.barh(names, [abs(c) for c in changes], color='green', alpha=0.7)
                ax2.set_title('Performance Improvements')
                ax2.set_xlabel('Performance Improvement (%)')
                ax2.grid(True, alpha=0.3)
                
                # Add value labels
                for i, (name, change, current, previous) in enumerate(improvements):
                    ax2.text(abs(change) + 0.5, i, f'{abs(change):.1f}%', 
                            va='center', ha='left')
            else:
                ax2.text(0.5, 0.5, 'No improvements detected', 
                        ha='center', va='center', transform=ax2.transAxes)
                ax2.set_title('Performance Improvements')
            
            plt.tight_layout()
            plt.savefig(self.history_dir / 'regression_analysis.png', dpi=300, bbox_inches='tight')
            plt.close()
    
    def generate_historical_report(self, history):
        """Generate a comprehensive historical report."""
        if not history.get('runs'):
            return
            
        report_lines = ["# Benchmark Historical Analysis\n"]
        report_lines.append(f"**Generated:** {datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S UTC')}\n")
        report_lines.append(f"**Total Runs:** {len(history['runs'])}\n")
        
        if len(history['runs']) >= 2:
            latest_run = history['runs'][-1]
            previous_run = history['runs'][-2]
            
            report_lines.append("## Recent Changes\n")
            report_lines.append(f"- **Latest Run:** {latest_run['timestamp'][:19]} (commit: {latest_run['commit']})")
            report_lines.append(f"- **Previous Run:** {previous_run['timestamp'][:19]} (commit: {previous_run['commit']})")
            
            # Calculate overall trend
            total_current = sum(r['ns_per_op'] for r in latest_run['results'].values())
#             total_previous = sum(r['ns_per_op'] for r in previous_run['results'].values()
#                                if r in latest_run['results'])
#
            if total_previous > 0:
                overall_change = ((total_current - total_previous) / total_previous) * 100
                trend_emoji = "📈" if overall_change > 0 else "📉"
                report_lines.append(f"- **Overall Trend:** {trend_emoji} {overall_change:+.1f}%")
        
        report_lines.append("\n## Generated Charts\n")
        report_lines.append("- [Performance Trends](performance_trends.png)")
        report_lines.append("- [Regression Analysis](regression_analysis.png)")
        
        with open(self.history_dir / "historical_report.md", "w") as f:
            f.write("\n".join(report_lines))

def main():
    # Load current benchmark results
    current_results_file = "benchmark-results/benchmark_data.json"
    
    if not os.path.exists(current_results_file):
        print("No current benchmark results found!")
        return
        
    with open(current_results_file, 'r') as f:
        current_results = json.load(f)
    
    print("Updating benchmark history...")
    tracker = BenchmarkHistoryTracker()
    
    # Update history with current results
    history = tracker.update_history(current_results)
    
    print(f"History now contains {len(history['runs'])} runs")
    
    # Generate historical analysis
    print("Generating historical analysis...")
    tracker.plot_performance_trends(history)
    tracker.plot_regression_analysis(history)
    tracker.generate_historical_report(history)
    
    print("✅ Historical analysis complete!")
    print(f"📊 Historical charts saved to: benchmark-results/history/")

if __name__ == "__main__":
    import numpy as np  # Import numpy for regression analysis
    main()