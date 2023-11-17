import React from 'react'
import './index.css'
import Sidebar from './components/sidebar'
import Docs from './components/docs'

function App() {

  return (
    <div className='appWrapper'>
      <Sidebar />
      <Docs />
    </div>
  )
}

export default App
