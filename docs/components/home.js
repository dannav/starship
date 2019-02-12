import React from 'react'
import { connect } from 'react-redux'
import Link from 'next/link'

class Home extends React.Component {
  state = {termTab: 'search'}

  constructor(props) {
    super(props)

    this.setTermTab = this.setTermTab.bind(this)
  }

  setTermTab(tab) {
    this.setState(old => {
      return {
        ...old,
        termTab: tab
      }
    })
  }

  render() {
    return (
      <div className="container">
        <p>
          Hi! I'm Danny Navarro. I'm a software engineer living in Florida who is currently working at <Link prefetch href="https://www.ardanlabs.com/"><a>Ardan Labs</a></Link>.
          You can check out what I work on by heading over to my <Link prefetch href="https://github.com/dannav"><a>Github</a></Link>, or chat with me on <Link prefetch href="https://twitter.com/danny_nav"><a>Twitter</a></Link>.
        </p>
      </div>
    )
  }
}

// filter redux state to props for this page
const mapStateToProps = ({ posts }) => ({ posts })

// connect redux state to component
export default connect(
  mapStateToProps,
  null
)(Home)