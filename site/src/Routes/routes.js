import React from 'react'
import {BrowserRouter, Route} from 'react-router-dom'
import {Greeting} from '../Greeting'
import Leader from '../Greeting/Leader'

class Routes extends React.Component {
  constructor(props){
    super(props)
  }
  render(){
    return (
    <BrowserRouter>
     <App />
    </BrowserRouter>
    )
  }
}

class App extends React.Component {
  constructor(props){
    super(props)
  }
  render(){
    return(
      <div>
        <Route exact path="/" component={Greeting} />
        <Route path="/leader" component={Leader}/>
      </div>
    )
  }
}

export default Routes