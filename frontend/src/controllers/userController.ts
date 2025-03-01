// src/controllers/userController.ts
import { Request, Response } from 'express';
import axios from 'axios';

// The Go backend API endpoint
const GO_BACKEND_URL = 'http://localhost:8080/LoadData'; // Go API endpoint

// Controller function to fetch data from the Go API and return it to the client
export const getLoadData = async (req: Request, res: Response) => {
  try {
    // Fetch data from the Go API
    const response = await axios.get(GO_BACKEND_URL);

    // Return the data as a JSON response to the client
    res.json(response.data);
  } catch (error) {
    console.error('Error fetching data from Go API:', error);
    res.status(500).json({ message: 'Error fetching data from Go API' });
  }
};
