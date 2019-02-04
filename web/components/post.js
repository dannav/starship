import React from 'react'
import { connect } from 'react-redux'
import marked from 'marked'

const rawMarkup = html => {
  return {
    __html: marked(html)
  }
}

class Post extends React.Component {
  render() {
    const post = this.props.posts.filter(p => p.id === parseInt(this.props.id, 10))[0]

    return (
      <>
        <h1>{post.title}</h1>
        <article dangerouslySetInnerHTML={rawMarkup(post.body)} />
      </>
    )
  }
}

export default connect(
  state => state
)(Post)