import http from 'k6/http';
import { check } from 'k6';

export default function () {
  const url = 'http://localhost:8080/pingpong';

  // Define the request body as an object
  const requestBody = {
    message: 'Ping'
  };

  // Run the request 100 times
 
    // Send a POST request with the JSON body
    const response = http.post(url, JSON.stringify(requestBody));

    check(response, {        'status is 200': (r) => r.status === 200,    });
  
}