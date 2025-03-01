// src/routes/userRoutes.ts
import express, { Request, Response } from 'express';
import { getLoadData } from '../controllers/userController';

const router = express.Router();

// Proxy the GET /users route to fetch data from the Go backend
router.get('/LoadData', getLoadData);

export default router;
