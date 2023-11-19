import React from 'react'
import Sidebar from './components/sidebar'
import Docs from './components/docs'
import { Toaster } from 'react-hot-toast'
import './index.css'

function App() {

  return (
    <div className='appWrapper'>
      <Toaster />
      <Sidebar />
      <Docs />
    </div>
  )
}

export default App
