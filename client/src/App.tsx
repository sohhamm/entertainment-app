import type {Component} from 'solid-js'
import '../src/styles/global.css'
import {Routes, Route} from '@solidjs/router'
import Home from './pages/home/Home'

const App: Component = () => {
  return (
    <Routes>
      <Route path='/' component={Home} />
    </Routes>
  )
}

export default App
