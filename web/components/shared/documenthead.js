import React from 'react'
import Head from 'next/head'

export default class DocumentHead extends React.Component {
  render() {
    return (
      <div>
          <link rel="manifest" href="static/manifest.json" />
          <meta charSet="utf-8" />
          <meta name="robots" content="index, follow" />
          <meta name="author" content="Starship" />
          <meta name="viewport" content="initial-scale=1.0, width=device-width" />
          <meta name="mobile-web-app-capable" content="yes" />
          <meta name="apple-mobile-web-app-capable" content="yes" />
          <meta name="application-name" content="Application" />
          <meta name="apple-mobile-web-app-title" content="Application" />
          <meta name="theme-color" content="#00b6b2" />
          <meta name="msapplication-navbutton-color" content="#00b6b2" />
          <meta name="apple-mobile-web-app-status-bar-style" content="black-translucent" />
          <meta name="msapplication-starturl" content="/" />
          <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
          <link rel="icon" type="image/png" sizes="192x192" href="static/icons/icon.png" />
          <link rel="apple-touch-icon" type="image/png" sizes="192x192" href="static/icons/icon.png" />
          <style jsx global>{`
          .btn.blue {
            background: #4a7eb3;
            color: #fff;
          }

          .btn.red {
            background: #f25c77;
            color: #fff !important;
          }

          .btn.green {
            background: #58a68d;
            color: #fff !important;
          }

          .btn:hover {
            cursor: pointer;
          }

          a, a:hover, a:visited, a:focus {
            color: #042B76;
          }

          body, html {
            margin: 0;
            padding: 0;
            height: 100%;
          }

          body {
            padding: 0 20px;
          }

          html {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
            font-size: 116%;
            color: #122239;
            -webkit-font-smoothing: antialiased;
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
            margin-bottom: 0;
            line-height: 1.5rem;
          }
          h1 {
            font-size: 2.121rem;
          }
          h2 {
            font-size: 1.414rem;
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
            margin-top: 1.5rem;
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
        `}</style>
      </div>
    )
  }
}