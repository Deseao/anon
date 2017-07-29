import React from 'react'
import {Link} from 'react-router-dom'

class Greeting extends React.Component {
  constructor(props){
    super(props)
  }
  render() {
    return (
      <div>
        <p>Welcome to Anon. Are you a Group Leader or a Participant?</p>
        <Link to="/leader">Leader</Link>
        <Link to="/participant">Participant</Link>
      </div>
    )
  }
}

Greeting.displayName = "Greeting"

export default Greeting