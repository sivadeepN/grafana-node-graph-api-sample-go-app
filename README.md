# Grafana Node Graph Sample App (Go)

<img width="751" alt="Screenshot 2023-09-30 at 10 55 21 AM" src="https://github.com/sivadeepN/grafana-node-graph-api-sample-go-app/assets/22883726/5d7e6026-be27-40c5-942c-ad82d9a0e946">



## Overview
This is a simple HTTP server written in Go that serves graph data and fields for a node graph. The server provides two main endpoints:

- `/api/graph/fields`: Returns the fields of the nodes and edges in the graph.
- `/api/graph/data`: Returns the data of the nodes and edges in the graph.

Additionally, there is a health check endpoint:

- `/api/health`: Returns the health status of the server.

## Usage
1. Fetch Graph Fields
     - Endpoint: `/api/graph/fields`
     - Response: 
     ```json
     {
         "nodes_fields": [
             {"field_name": "id", "type": "string"},
             {"field_name": "title", "type": "string"},
             {"field_name": "subTitle", "type": "string"},
             {"field_name": "mainStat", "type": "string"},
             {"field_name": "secondaryStat", "type": "number"},
             {"field_name": "arc__failed", "type": "number", "color": "red", "displayName": "Failed"},
             {"field_name": "arc__passed", "type": "number", "color": "green", "displayName": "Passed"},
             {"field_name": "detail__role", "type": "string", "displayName": "Role"}
         ],
         "edges_fields": [
             {"field_name": "id", "type": "string"},
             {"field_name": "source", "type": "string"},
             {"field_name": "target", "type": "string"},
             {"field_name": "mainStat", "type": "number"}
         ]
     }
     ```
2. Fetch Graph Data
     - Endpoint: `/api/graph/data`
     - Response:
     ```json
     {
         "nodes": [
             {"id": "1", "title": "Service1", "subTitle": "instance:#2", "detail__role": "load", "arc__failed": 0.7, "arc__passed": 0.3, "mainStat": "qaz"},
             {"id": "2", "title": "Service2", "subTitle": "instance:#2", "detail__role": "transform", "arc__failed": 0.5, "arc__passed": 0.5, "mainStat": "qaz"},
             {"id": "3", "title": "Service3", "subTitle": "instance:#3", "detail__role": "extract", "arc__failed": 0.3, "arc__passed": 0.7, "mainStat": "qaz"},
             {"id": "4", "title": "Service3", "subTitle": "instance:#1", "detail__role": "transform", "arc__failed": 0.5, "arc__passed": 0.5, "mainStat": "qaz"},
             {"id": "5", "title": "Service4", "subTitle": "instance:#5", "detail__role": "transform", "arc__failed": 0.5, "arc__passed": 0.5, "mainStat": "qaz"}
         ],
         "edges": [
             {"id": "1", "source": "1", "target": "2", "mainStat": 53},
             {"id": "2", "source": "2", "target": "3", "mainStat": 53},
             {"id": "2", "source": "1", "target": "4", "mainStat": 5},
             {"id": "3", "source": "3", "target": "5", "mainStat": 70},
             {"id": "4", "source": "2", "target": "5", "mainStat": 100}
         ]
     }
     ```
3. Health Check
     - Endpoint: `/api/health`
     - Response:
     ```json
     {
         "message": "App is up",
         "status": "ok"
     }
     ```
     

## Running the Server
1. Clone the repository: `git clone https://github.com/yourusername/node-graph-sample-app-go.git`
2. Navigate to the project directory: `cd node-graph-sample-app-go`
3. Run the Go application: `go run main.go`
4. The server will start on `http://localhost:5000`.
5. URL : http://node-graph-sample-app-go.tempo-test.svc.cluster.local:5000
<img width="1134" alt="Screenshot 2023-09-30 at 10 55 53 AM" src="https://github.com/sivadeepN/grafana-node-graph-api-sample-go-app/assets/22883726/8589404b-4595-4959-9934-cfb6411c3886">

Feel free to explore and integrate this simple Go server into your projects! If you have any questions or encounter issues, please don't hesitate to reach out.
