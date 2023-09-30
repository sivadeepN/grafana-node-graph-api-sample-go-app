// Package main provides a simple HTTP server that serves a graph data and fields.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// fetchGraphFields is a handler function that returns the fields of the graph.
func fetchGraphFields(w http.ResponseWriter, r *http.Request) {
	// nodesFields and edgesFields are the fields of the nodes and edges of the graph.
	nodesFields := []map[string]interface{}{
		{"field_name": "id", "type": "string"},
		{"field_name": "title", "type": "string"},
		{"field_name": "subTitle", "type": "string"},
		{"field_name": "mainStat", "type": "string"},
		{"field_name": "secondaryStat", "type": "number"},
		{"field_name": "arc__failed", "type": "number", "color": "red", "displayName": "Failed"},
		{"field_name": "arc__passed", "type": "number", "color": "green", "displayName": "Passed"},
		{"field_name": "detail__role", "type": "string", "displayName": "Role"},
	}

	edgesFields := []map[string]interface{}{
		{"field_name": "id", "type": "string"},
		{"field_name": "source", "type": "string"},
		{"field_name": "target", "type": "string"},
		{"field_name": "mainStat", "type": "number"},
	}

	// result is a map containing the nodes and edges fields.
	result := map[string]interface{}{
		"nodes_fields": nodesFields,
		"edges_fields": edgesFields,
	}

	// Set the content type of the response to application/json and encode the result as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// fetchGraphData is a handler function that returns the data of the graph.
func fetchGraphData(w http.ResponseWriter, r *http.Request) {
	// nodes and edges are the data of the nodes and edges of the graph.
	nodes := []map[string]interface{}{
		{"id": "1", "title": "Service1", "subTitle": "instance:#2", "detail__role": "load", "arc__failed": 0.7, "arc__passed": 0.3, "mainStat": "qaz"},
		{"id": "2", "title": "Service2", "subTitle": "instance:#2", "detail__role": "transform", "arc__failed": 0.5, "arc__passed": 0.5, "mainStat": "qaz"},
		{"id": "3", "title": "Service3", "subTitle": "instance:#3", "detail__role": "extract", "arc__failed": 0.3, "arc__passed": 0.7, "mainStat": "qaz"},
		{"id": "4", "title": "Service3", "subTitle": "instance:#1", "detail__role": "transform", "arc__failed": 0.5, "arc__passed": 0.5, "mainStat": "qaz"},
		{"id": "5", "title": "Service4", "subTitle": "instance:#5", "detail__role": "transform", "arc__failed": 0.5, "arc__passed": 0.5, "mainStat": "qaz"},
	}

	edges := []map[string]interface{}{
		{"id": "1", "source": "1", "target": "2", "mainStat": 53},
		{"id": "2", "source": "2", "target": "3", "mainStat": 53},
		{"id": "2", "source": "1", "target": "4", "mainStat": 5},
		{"id": "3", "source": "3", "target": "5", "mainStat": 70},
		{"id": "4", "source": "2", "target": "5", "mainStat": 100},
	}

	// result is a map containing the nodes and edges data.
	result := map[string]interface{}{
		"nodes": nodes,
		"edges": edges,
	}

	// Set the content type of the response to application/json and encode the result as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// checkHealth is a handler function that returns the health status of the server.
func checkHealth(w http.ResponseWriter, r *http.Request) {
	// response is a map containing the health status of the server.
	response := map[string]string{"message": "App is up", "status": "ok"}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	// Set the content type of the response to application/json and write the response as JSON.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func main() {
	// Register the handler functions for the HTTP server.
	http.HandleFunc("/api/graph/fields", func(w http.ResponseWriter, r *http.Request) {
		fetchGraphFields(w, r)
	})

	http.HandleFunc("/api/graph/data", func(w http.ResponseWriter, r *http.Request) {
		fetchGraphData(w, r)
	})

	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		checkHealth(w, r)
	})

	// Start the HTTP server on port 5000.
	fmt.Println("Server is running on :5000...")
	http.ListenAndServe(":5000", nil)
}
