import React from 'react'

export const HeadContent = () => {
  return (
    <>
      <meta charSet="utf-8" />
      <meta name="robots" content="index, follow" />
      <meta name="author" content="Danny Navarro" />
      <meta name="viewport" content="initial-scale=1.0, width=device-width" />
      <meta name="mobile-web-app-capable" content="yes" />
      <meta name="apple-mobile-web-app-capable" content="yes" />
      <meta name="application-name" content="Application" />
      <meta name="apple-mobile-web-app-title" content="Application" />
      <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />

      <style jsx global>{`
      html {
        font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
        font-size: 116%;
        color: #122239;
        -webkit-font-smoothing: antialiased;
      }

      body {
        padding: 15px;
      }

      p {
        line-height: 1.5rem;
        margin-top: 0.75rem;
        margin-bottom: 0;
      }

      ul,
      ol {
        margin-top: 0.75rem;
        margin-bottom: 1.5rem;
      }

      ul li,
      ol li {
        line-height: 1.5rem;
      }

      ul ul,
      ol ul,
      ul ol,
      ol ol {
        margin-top: 0;
        margin-bottom: 0;
      }

      ul {
        margin: 0;
        padding: 0;
        list-style: none;
      }

      li {
        padding: 0;
      }

      blockquote {
        line-height: 1.5rem;
        margin-top: 0.75rem;
        margin-bottom: 0.75rem;
      }

      h1,
      h2,
      h3,
      h4,
      h5,
      h6 {
        margin-top: 0.75rem;
        margin-bottom: 0.75rem;
        line-height: 1.5rem;
      }

      h1 {
        font-size: 2.121rem;
      }

      h2 {
        font-size: 1.114rem;
      }

      h3 {
        font-size: 0.707rem;
      }

      h4 {
        font-size: 0.707rem;
      }

      h5 {
        font-size: 0.4713333333333333rem;
      }

      h6 {
        font-size: 0.3535rem;
      }

      table {
        margin: 1.5rem 0;
        border-spacing: 0px;
        border-collapse: collapse;
      }

      table td,
      table th {
        padding: 0;
        line-height: 33px;
      }

      code {
        vertical-align: bottom;
      }

      .lead {
        font-size: 1.414rem;
      }

      .hug {
        margin-top: 0;
      }

      a, a:hover, a:visited, a:focus {
        color: #122239;
        text-decoration: none;
      }

      a:visited {
        color: black;
      }

      a:hover {
        border-bottom: 1px solid #122239;
      }

      code {
        background: #eaeaea;
        color: #122239;
        padding: 3px 5px;
      }

      hr {
        outline: none;
        border: 1px solid #eaeaea;
      }

      table {
        white-space: nowrap;
      }
    `}</style>
    </>
  )
}
export default class DocumentHead extends React.Component {
  render() {
    return (
      <HeadContent />
    )
  }
}