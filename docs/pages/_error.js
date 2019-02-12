import React from 'react'
import { connect } from 'react-redux'
import Layout from '../components/shared/layout'

class NotFound extends React.Component {
  static async getInitialProps ({ store, isServer }) {
    return { isServer }
  }

  render() {
    return (
      <Layout title='Starship'>
        <div className="container">
          <p>
            The page you were looking for could not be found.
          </p>
        </div>
      </Layout>
    )
  }
}

export default connect(state => state)(NotFound)