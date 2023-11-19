import React from 'react'
import Sidebar from './components/sidebar'
import Docs from './components/docs'
import { Toaster } from 'react-hot-toast'
import './index.css'
import Nouns from './components/nouns'

function App() {

  return (
    <div className='appWrapper'>
      <Toaster />
      <Nouns />
      <Sidebar />
      <Docs />
    </div>
  )
}

export default App
