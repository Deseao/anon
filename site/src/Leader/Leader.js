import React from 'react'
import {Link, Route} from 'react-router-dom'
import CreateGroup from './CreateGroup'

const Leader = ({match}) => {
  return (
    <div>
      <Link to="/leader/createGroup">Create New Group</Link>
      <Link to="/sendMessage">Send a Message to an Existing Group</Link>
    </div>
  )
}

export default Leader