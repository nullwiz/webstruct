import 'reactflow/dist/style.css'
import { useCallback } from 'react'
import ReactFlow, { ReactFlowProvider, useReactFlow } from 'reactflow'

import defaultNodes from './nodes.js'
import defaultEdges from './edges.js'

import './button.css'

const edgeOptions = {
  animated: true,
  style: {
    stroke: 'white',
  },
}

const connectionLineStyle = { stroke: 'white' }

let nodeId = 0

function Flow() {
  const reactFlowInstance = useReactFlow()
  const onClick = useCallback(() => {
    const id = `${++nodeId}`
    const newNode = {
      id,
      position: {
        x: Math.random() * 500,
        y: Math.random() * 500,
      },
      data: {
        label: `Node ${id}`,
      },
    }
    reactFlowInstance.addNodes(newNode)
  }, [])

  const proOptions = { hideAttribution: true }

  return (
    <>
      <ReactFlow
        defaultNodes={defaultNodes}
        defaultEdges={defaultEdges}
        defaultEdgeOptions={edgeOptions}
        fitView
        style={{
          backgroundColor: '#2a303c',
        }}
        connectionLineStyle={connectionLineStyle}
        proOptions={proOptions}
      />
    </>
  )
}

export default function FlowInstance() {
  return (
    <ReactFlowProvider>
      <Flow />
    </ReactFlowProvider>
  )
}
