# Testing Guide

## Performance Testing

The following metrics are tested across various environments and inputs to ensure an optimal and efficient user experience:
- Framerate (FPS) 
- Memory Usage (MB)

Performance tests are located in `tests/performance`. There is a helper file `simulation.go` that simulates the application. Other files with the pattern `*_test.go` contain the actual performance tests and metric collection logic.

There are automated performance tests that check the above metrics meet a certain threshold. There is also a export flag in performance testing that exports the performance metrics to a JSON file as a time series data set for further analysis.

To run the automated tests, use the following command (from root). 
```bash
go test -v tests/performance/*.go
```

To export the data to JSON, add a export flag to the command. Then, use the python script in `scripts/` to graph the data.

```bash
go test -v tests/performance/*.go -export
python scripts/graph_performance.py 
```

## Unit Testing 

Unit tests are written to explicitly test expected behavior when the application is in a certain state. The unit tests are not comprehensive and focus mainly on core logic that is critical to the functionality of the application. Unit tests should also largely focus on edge cases to ensure the application is robust and can handle various unintended states.

To run the unit tests, use the following command (from project root):

```bash
go test -v tests/unit/*.go
```

