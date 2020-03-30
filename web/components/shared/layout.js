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
          <h1>{props.title}</h1>
          <nav className="menu">
            <ul>
              <li>
                <Link prefetch href="/">
                  <a>Home</a>
                </Link>
              </li>
              <li>
                <Link prefetch href="/contact">
                  <a>Use Cases</a>
                </Link>
              </li>
              <li>
                <Link prefetch href="/guide">
                  <a>Documentation</a>
                </Link>
              </li>
            </ul>
          </nav>
          <nav className="header-actions">
            <ul>
              <li>
                <Link prefetch href="/">
                  <a className="dwnld">Download</a>
                </Link>
              </li>
              <li>
                <Link prefetch href="/contact">
                  <a className="btn">Login / Sign up</a>
                </Link>
              </li>
            </ul>
          </nav>
        </header>
        <main>
          {
            childrenWithProps(props.children, props)
          }
        </main>
        <footer>
          <div className="container">
            <nav>
              <ul>
                <li>Privacy &amp; Terms</li>
                <li>Contact Us</li>
              </ul>
            </nav>
          </div>
        </footer>
        <style global jsx>{`
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

          header.container {
            max-width: 1090px;
          }

          .container {
            max-width: 940px;
            margin: 0 auto;
          }

          .btn {
            padding: 8px 20px;
            text-align: center;
            border-radius: 7px;
            background: #122239;
            color: #fff !important;
            font-size: 90%;
            border: none !important;
          }

          header .btn {
            padding: 5px 15px;
          }

          .btn:hover {
            border: none;
          }

          header {
            width: 100%;
            display: grid;
            grid-template-columns: .75fr 1fr .75fr;
            align-items: center;
            height: 100px;
            box-sizing: border-box;
          }

          header h1 {
            margin: 0;
            font-size: 1.414rem;
            position: relative;
            top: -2px;
          }

          header .menu {
            justify-self: center;
          }

          header .header-actions {
            justify-self: end;
          }

          header .dwnld {
            color: #707d8e !important;
          }

          header .dwnld:hover {
            border: 0;
            color: #122239 !important;
          }

          header a, header a:hover, header a:visited, header a:focus {
            color: #122239;
            text-decoration: none;
          }

          header a:visited {
            color: black;
          }

          header a:hover {
            border-bottom: 1px solid #122239;
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
        `}</style>
      </>
    )
  }
)