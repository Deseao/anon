import React from 'react'
import {apiUrl} from '../../config'

class Participant extends React.Component {
  constructor(props) {
    super(props)
    this.state ={
      code:'',
      email:'',
      phone:'',
      isSubscribed: false
    }

    this.subscribe = this.subscribe.bind(this)
  }

  setGroup(code){
    this.setState({code})
  }

  setEmail(email){
    this.setState({email})
  }
  setPhone(phone){
    this.setState({phone})
  }
  subscribe(){
    fetch(apiUrl + "/signup", {method:'POST',
      body: JSON.stringify({code: this.state.code, email: this.state.email, phone: this.state.phone})
    })
      .then(response => {
        if (!response.ok) {
          console.log("Error: ", response.ok)
          return
        }
        this.setState({isSubscribed:true})
      })
  }

  isSubscribed(){
    if (!this.state.isSubscribed) {return null }
    return <p>Subscribed!</p>
  }

  render() {
    return (
    <div>
      <p> Subscribe to get notified! </p>
      Group Code: <input type="text" onChange={(e) => this.setGroup(e.target.value.trim())} />
      <br/>
      Email: <input type="text" onChange={(e) => this.setEmail(e.target.value.trim())} />
      <br />
      Phone: <input type="text" onChange={(e) => this.setPhone(e.target.value.trim())} />
      <br />
      <button type="button" onClick={this.subscribe}>Subscribe</button>
      {this.isSubscribed()}
    </div>
    )
  }

}

Participant.displayName = "Participant"

export default Participant