import React from 'react'
import {apiUrl} from '../../config'

class CreateGroup extends React.Component {
  constructor(props) {
    super(props)

    this.state = {
      groupCode:''
    }

    this.createGroup = this.createGroup.bind(this)
  }
  createGroup(){
    fetch(apiUrl + "/create", {method:'POST'})
      .then(response => response.json())
      .then(json => {
        this.setState({groupCode:json.code})
      })
  }
  displayCode(){
    if (this.state.groupCode === '') { return null}
    return <p>Your new group code is {this.state.groupCode}</p>
  }

  render(){
    return(
      <div>
        <button type="button" onClick={this.createGroup}>Click to Create Group</button>
        {this.displayCode()}
      </div>
    )
  }
}

CreateGroup.displayName = "CreateGroup"

export default CreateGroup