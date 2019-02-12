import React from 'react'
import { connect } from 'react-redux'
import Layout from '../components/shared/layout'
import { getContent } from '../redux/store'
import Head from 'next/head'
import { HeadContent } from '../components/shared/documenthead'
const md = require('markdown-it')({html: true})

const Body = props => {
  return (
    <div className="container" dangerouslySetInnerHTML={{__html: md.render(props.content)}} />
  )
}

class Doc extends React.Component {
  static async getInitialProps ({ res, store, isServer, query: { folder, page } }) {
    await store.dispatch(getContent(isServer, res, folder, page))

    return {
      isServer,
      folder,
      page,
    }
  }

  render() {
    const pageTitles = this.props.page.split('-').map(t => t.capitalize()).join(' ')
    const folderTitles = this.props.folder.split('-').map(t => {
      if (t == 'api') {
        return 'API'
      }

      return t.capitalize()
    }).join(' ')

    const seoTitle = `${folderTitles} - ${pageTitles}`

    return (
      <>
      <Head>
        <title>{seoTitle}</title>
        <HeadContent />
      </Head>
      <Layout title="Starship">
        <Body content={this.props.content} />
      </Layout>
      </>
    )
  }
}

export default connect(state => state)(Doc)