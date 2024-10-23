#!/bin/bash
 
cd frontend && npm start &
cd backend && go run main.go &
wait