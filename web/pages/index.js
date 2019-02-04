import React from 'react'
import { connect } from 'react-redux'
import Home from '../components/homepage'
import Layout from '../components/shared/layout'
import { getPosts } from '../redux/store'

class Homepage extends React.Component {
  static async getInitialProps ({ store, isServer }) {
    await store.dispatch(getPosts(isServer))
    return { isServer }
  }

  render() {
    return (
      <Layout title='Starship'>
        <Home />
      </Layout>
    )
  }
}

export default connect(state => state)(Homepage)