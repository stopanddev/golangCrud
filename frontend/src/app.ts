import express, {Request, Response} from 'express';
import userRoutes from './routes/userRoutes';
import path from 'path';


const app = express();

// Serve static files form the views folder
app.use(express.static(path.join(__dirname, 'views')));
// Middleware to parse JSON request body
app.use(express.json());
app.use('/', userRoutes);

//Define a route for home page
app.get('/', (req: Request, res: Response) => {
    res.sendFile(path.join(__dirname, 'views', '../../index.html'))
})
// Use the userRoutes for handling user related requests

export default app;
