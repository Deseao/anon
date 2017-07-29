import React from 'react'
import {apiUrl} from '../../config'

class SendMessage extends React.Component {
  constructor(props) {
    super(props)

    this.state = {
      code:'',
      message:'',
      isSent: false
    }

    this.sendMessage = this.sendMessage.bind(this)
  }

  setGroup(code) {
  this.setState({code})
  }

  setMessage(message) {
  this.setState({message})
  }

  sendMessage(){
    fetch(apiUrl + "/send", {method:"POST",
    body: JSON.stringify({code: this.state.code, message: this.state.message})
    }).then((response) => {
      if (!response.ok) {
        console.log("Response error: ", response.ok)
      }
      this.setState({isSent:true})
    })
  }

  messageSent(){
    if(!this.state.isSent) { return false}
    return <p>Message Sent!</p>
  }

  render(){
    return (
      <div>
        <p>Send a message to your group</p>
        Group Code: <input type="text" onChange={(e) => this.setGroup(e.target.value.trim())} />
        <br />
        Message:
        <br />
        <textarea onChange={(e)=> this.setMessage(e.target.value)} defaultValue="Enter your message here." />
        <br />
        <button type="button" onClick={this.sendMessage}>Send!</button>
        {this.messageSent()}
      </div>
    )
  }
}

SendMessage.displayName = "SendMessage"

export default SendMessage