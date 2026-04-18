import json
import matplotlib.pyplot as plt
from dateutil import parser
from dataclasses import dataclass
from typing import Any
from pathlib import Path


@dataclass
class GraphData:
    timestamp: list[float]
    value: list[float]


@dataclass
class Metadata:
    metric: str
    title: str


@dataclass
class PerformanceData:
    metadata: Metadata
    data_points: GraphData


def read_json(json_file: Path) -> dict[str, Any]:
    if not json_file.exists():
        raise FileNotFoundError(f"File {json_file} does not exist.")
    with open(json_file, "r") as f:
        return json.load(f)


def serialize_data(data: dict[str, Any]) -> PerformanceData:
    graph_data: GraphData = GraphData(timestamp=[], value=[])
    start_time = -1
    for entry in data.get("data", []):
        time = parser.parse(entry.get("timestamp"))
        start_time = start_time if start_time != -1 else time
        normalized_time = (time - start_time).total_seconds()
        graph_data.timestamp.append(normalized_time)
        graph_data.value.append(entry.get("value", 0.0))
    metadata_json = data.get("metadata", {})
    metadata = Metadata(
        metric=metadata_json.get("metric", "unknown"),
        title=metadata_json.get("title", "unknown"),
    )
    return PerformanceData(metadata=metadata, data_points=graph_data)


if __name__ == "__main__":
    results_dir = Path("tests/performance/test_results").glob("*.json")
    for json_file in results_dir:
        data = read_json(json_file)
        performance_data = serialize_data(data)
        plt.figure(figsize=(10, 5))
        plt.plot(
            performance_data.data_points.timestamp,
            performance_data.data_points.value,
            marker="x",
            color="red",
        )
        plt.title(performance_data.metadata.title)
        plt.xlabel("Time (seconds)")
        plt.ylabel(performance_data.metadata.metric)
        plt.grid(True)
        output_file = json_file.with_suffix(".png")
        plt.savefig(output_file)
        print(f"Plot saved to {output_file}")
