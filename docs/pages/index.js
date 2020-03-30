import React from 'react'
import { connect } from 'react-redux'
import Layout from '../components/shared/layout'
import Head from 'next/head'
import { HeadContent } from '../components/shared/documenthead'
import { getContent } from '../redux/store'
const md = require('markdown-it')({html: true})

const Body = props => {
  return (
    <div className="container homepage" dangerouslySetInnerHTML={{__html: md.render(props.content)}} />
  )
}

class Homepage extends React.Component {
  static async getInitialProps ({ res, store, isServer, query: { folder, page } }) {
    await store.dispatch(getContent(isServer, res, 'home', 'index'))

    return {
      isServer,
      folder: 'home',
      page: 'index',
    }
  }

  render() {
    return (
      <>
      <Head>
        <title>Starship Documentation</title>
        <HeadContent />
      </Head>
      <Layout title="Starship">
        <Body content={this.props.content} />
      </Layout>
      </>
    )
  }
}

export default connect(state => state)(Homepage)