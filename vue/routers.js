import Graph from './graph.html'
import Home from './index.html'

const routers = [
    {
        path: '/graph',
        name: 'graph',
        component: Graph
    },
    {
         path: '/',
         component: Home
    },
]
export default routers
