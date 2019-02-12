import { connect } from 'react-redux'
import Link from 'next/link'

const childrenWithProps = (children, props) => React.Children.map(children, child =>
  React.cloneElement(child, props)
);

// connect redux state to component
export default connect(state => state)(
  props => {
    return (
      <>
        <header className="container">
          <Link prefetch href="/">
            <a className="logo">
              <h1>{props.title}</h1>
            </a>
          </Link>
          <div className="header-links">
            <Link prefetch href="https://github.com/starship-fyi/starship">
              <a>
                View on Github
              </a>
            </Link>
          </div>
        </header>
        <main>
          <aside>
            <ul>
              <li className="section">
                <h2>Getting Started</h2>
              </li>
              <li>
                  <Link prefetch href="/doc?folder=getting-started&page=introduction" as="/getting-started/introduction">
                    <a>Introduction</a>
                  </Link>
              </li>
              <li>
                  <Link prefetch href="/doc?folder=getting-started&page=architecture-overview" as="/getting-started/architecture-overview">
                    <a>Architecture Overview</a>
                  </Link>
              </li>
              <li>
                  <Link prefetch href="/doc?folder=getting-started&page=running-starship" as="/getting-started/running-starship">
                    <a>Running Starship</a>
                  </Link>
              </li>
              <li className="section">
                <h2>API Reference</h2>
              </li>
              <li>
                  <Link prefetch href="/doc?folder=api-reference&page=index" as="/api-reference/index">
                    <a>Index</a>
                  </Link>
              </li>
              <li>
                  <Link prefetch href="/doc?folder=api-reference&page=search" as="/api-reference/search">
                    <a>Search</a>
                  </Link>
              </li>
              <li>
                  <Link prefetch href="/doc?folder=api-reference&page=download" as="/api-reference/download">
                    <a>Download</a>
                  </Link>
              </li>
              <li>
                  <Link prefetch href="/doc?folder=api-reference&page=exists" as="/api-reference/exists">
                    <a>Exists</a>
                  </Link>
              </li>
              <li>
                  <Link prefetch href="/doc?folder=api-reference&page=ready" as="/api-reference/ready">
                    <a>Ready</a>
                  </Link>
              </li>
              <li className="section">
                <h2>Starship CLI</h2>
              </li>
              <li>
                  <Link prefetch href="/doc?folder=getting-started&page=get-starship-cli" as="/getting-started/get-starship-cli">
                    <a>Get The Application</a>
                  </Link>
              </li>
              <li>
                  <Link prefetch href="/doc?folder=getting-started&page=using-the-starship-cli" as="/getting-started/using-the-starship-cli">
                    <a>Using The CLI</a>
                  </Link>
              </li>
              <li className="section">
                <h2>Misc</h2>
              </li>
              <li>
                  <Link prefetch href="/doc?folder=misc&page=contributing" as="/misc/contributing">
                    <a>Contributing</a>
                  </Link>
              </li>
              <li>
                  <Link prefetch href="/doc?folder=misc&page=code-of-conduct" as="/misc/code-of-conduct">
                    <a>Code of Conduct</a>
                  </Link>
              </li>
              <li style={{color: '#fff', background: '#122239', textAlign: 'center'}}>
                  <Link prefetch href="/doc?folder=misc&page=sponsorship" as="/misc/sponsorship">
                    <a style={{color: '#fff'}}>Sponsor Starship</a>
                  </Link>
              </li>
            </ul>
          </aside>
          <div className="main-content">
            {
              childrenWithProps(props.children, props)
            }
          </div>
        </main>
        <style global jsx>{`
          .main-content .container > *:first-child {
            margin-top: 0;
          }

          .main-content a {
            color: blue;
          }

          .main-content a:hover {
            border-color: blue;
          }

          .main-content img {
            height: 100%;
            margin: 2rem auto;
            display: block;
          }

          .section h2 {
            margin: 0;
            padding-bottom: .25rem;
            margin-top: .75rem;
          }

          aside ul .section:first-child h2 {
            margin-top: 0 !important;
          }

          main {
            display: grid;
            grid-template-columns: fit-content(300px) fit-content(740px);
            column-gap: 2rem;
            padding-top: 0.75rem;
          }

          footer {
            font-weight: bold;
            color: #122239;
            padding: 1.5rem 0;
          }

          footer ul {
            list-style: none;
            margin: 0;
            padding: 0;
          }

          footer li {
            display: inline-block;
            font-size: 80%;
            margin-right: 1rem;
          }

          .container {
            max-width: 740px;
          }

          header {
            margin-bottom: 1rem;
            max-width: 960px !important;
          }

          header h1 {
            margin: 0;
            font-size: 1.414rem;
            position: relative;
            top: -2px;
          }

          header h1 {
            justify-self: start;
          }

          header ul {
            padding: 0;
            margin: 0;
            list-style: none;
          }

          header li {
            display: inline-block;
            padding-left: 20px;
          }

          header li:first-child {
            padding: 0;
          }

          .logo:hover {
            border: none;
          }

          .logo h1 {
            display: inline-block;
          }

          .header-links {
            display: inline-block;
            float: right;
          }

          table td {
            padding: 10px !important;
          }

          table {
            border: 1px solid #122239;
          }

          th {
            background: #122239;
            text-align: left !important;
            border-bottom: 1px solid #122239;
            padding: 5px 10px !important;
            color: #fff;
          }

          table td {
            border-left: 1px solid #122239;
            border-right: 1px solid #122239;
            border-bottom: 1px solid #122239;
          }

          table td:first-child {
              border-left: none;
          }

          table td:last-child {
              border-right: none;
          }

          pre {
            background: #eaeaea;
            padding: 10px;
            overflow: scroll;
          }

          pre code {
            padding: 0;
          }

          ul li ul {
            margin-left: 1rem;
          }
        `}</style>
      </>
    )
  }
)