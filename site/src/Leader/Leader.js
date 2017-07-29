import React from 'react'
import {Link} from 'react-router-dom'
const Leader = () => {
  return (
    <div>
      <Link to="/leader/createGroup">Create New Group</Link>
      <Link to="/leader/sendMessage">Send a Message to an Existing Group</Link>
    </div>
  )
}

export default Leader