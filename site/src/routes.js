import React from 'react'
import {BrowserRouter, Route} from 'react-router-dom'
import {Greeting} from './Greeting/index'
import {Leader, CreateGroup} from './Leader/index'
import {Participant} from './Participant/index'

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
        <Route exact path="/leader" component={Leader}/>
        <Route exact path="/leader/createGroup" component={CreateGroup} />
        <Route exact path="/participant" component={Participant} />
      </div>
    )
  }
}

export default Routes