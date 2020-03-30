import React from 'react'
import { connect } from 'react-redux'
import Contact from '../components/contact'
import Layout from '../components/shared/layout'

class ContactPage extends React.Component {
  render() {
    return (
      <Layout title="Send me an email I won't read">
        <Contact />
      </Layout>
    )
  }
}

export default connect(state => state)(ContactPage)