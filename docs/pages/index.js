import React from 'react'
import { connect } from 'react-redux'
import Home from '../components/home'
import Layout from '../components/shared/layout'
import Head from 'next/head'
import { HeadContent } from '../components/shared/documenthead'

class Homepage extends React.Component {
  static async getInitialProps ({ store, isServer }) {
    return { isServer }
  }

  render() {
    return (
      <>
        <Head>
          <title>Starship Documentation</title>
          <HeadContent />
        </Head>
        <Layout title='Starship'>
          <Home />
        </Layout>
      </>
    )
  }
}

export default connect(state => state)(Homepage)