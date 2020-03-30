import React from 'react'
import { connect } from 'react-redux'
import Post from '../components/post'
import Layout from '../components/shared/layout'
import { getPosts } from '../redux/store'

class PostPage extends React.Component {
  // getInitialProps is executed by nextjs on the server and passes the redux store, query params,
  // and a value determining if we are executing on the server or not
  static async getInitialProps ({ store, isServer, query: { id } }) {
    await store.dispatch(getPosts(isServer))
    return { isServer, postID: id }
  }

  render() {
    return (
      <Layout>
        <Post id={this.props.postID} />
      </Layout>
    )
  }
}

export default connect(state => state)(PostPage)