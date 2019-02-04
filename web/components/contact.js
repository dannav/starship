import React from 'react'
import { connect } from 'react-redux'

class Contact extends React.Component {
  state = {
    name: '',
    email: '',
  }

  handleChange(prop, e) {
    const value = e.target.value

    this.setState(_ => {
      let newState = {}
      newState[prop] = value
      return {
        ...this.state,
        ...newState
      }
    })
  }

  submit() {
    console.log(this.state)
  }

  render() {
    return (
      <>
        <div className="form">
          <label>
            Name:
            <input type="text" placeholder="name" value={this.state.name} onChange={this.handleChange.bind(this, 'name')} />
          </label>
          <label>
            Email:
            <input type="text" placeholder="email" value={this.state.email} onChange={this.handleChange.bind(this, 'email')} />
          </label>
          <textarea value={this.state.message} onChange={this.handleChange.bind(this, 'message')} />
          <button type="button" onClick={this.submit.bind(this)}>Submit</button>
        </div>

        <style jsx>{`
          label {
            display: block;
            margin-bottom: 10px;
          }

          label input {
            margin-left: 5px;
          }

          textarea {
            width: 100%;
            height: 100px;
          }
        `}</style>
      </>
    )
  }
}

export default connect(
  state => state
)(Contact)