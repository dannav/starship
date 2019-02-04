import React from 'react'
import { connect } from 'react-redux'
import Link from 'next/link'

class Home extends React.Component {
  state = {termTab: 'search'}

  constructor(props) {
    super(props)

    this.setTermTab = this.setTermTab.bind(this)
  }

  setTermTab(tab) {
    this.setState(old => {
      return {
        ...old,
        termTab: tab
      }
    })
  }

  render() {
    return (
      <>
        <div className="full-width">
          <div className="stars"></div>
          <div className="twinkling"></div>
          <div className="clouds"></div>

          <div className="hero">
            <div className="container">
              <h2>
                Keep Your Remote Team In Sync
              </h2>
              <p>
                {/* An effective team requires constant access to internal information. Provide them with the tools to find whatever they need, whenever they need it. */}
                Starship is a private search engine for you and your team. Make finding the information anyone on your team needs one command away.
              </p>
              <div className="hero-actions">
                <div>
                  <button type="button" className="btn green">Starship in 60 seconds</button>
                </div>
                <div>
                  <button type="button" className="btn red">Get started - It's free</button>
                </div>
              </div>
            </div>
          </div>
          <div className="container tab-section pull-front">
            <ul>
              <li className={this.state.termTab === 'search' ? 'active' : null}>
                <button type="button" onClick={this.setTermTab.bind(this, 'search')}>Search</button>
              </li>
              <li className={this.state.termTab === 'index' ? 'active' : null}>
                <button type="button" onClick={this.setTermTab.bind(this, 'index')}>Index</button>
              </li>
            </ul>
          </div>
        </div>
        <div className="showcase full-width">
          <div className="container">
            <div className="terminal" id="termynal" data-terminal>
              {
                (() => {
                  let url = '/static/search.svg' + "?" + Math.random()
                  if (this.state.termTab === 'search') {
                    return <img src={url} />
                  }

                  url = '/static/index.svg' + "?" + Math.random()
                  return <img src={url} />
                })()
              }
            </div>
          </div>
        </div>
        <div className="showcase full-width alt">
          <div className="container">
            <h2>Your Private Search Engine</h2>
            <p>
              Sifting through project documentation, company information, and your personal notes. Starship lets your team upload any document and search them with one command.
            </p>
            <div className="two-sided">
              <div className="placeholder"></div>
              <div className="description">
                <p>
                  Starship was built to make sharing documented information in remote teams easier.
                </p>
                <p>
                  Store PDF's, Word documents, HTML pages, Markdown, or just about any other text based file format.
                </p>
                <p>
                  We'll index the content and make it searchable so that you can share it with others and find all of your team's information later.
                </p>
              </div>
            </div>
          </div>
        </div>
        <div className="showcase full-width">
          <div className="container three">
            <div>
              <h2>AI-Powered</h2>
              <p className="hug">
                Starship uses AI to provide better results than your typical search.
              </p>
            </div>
            <div>
              <h2>Security Focused</h2>
              <p className="hug">All documents stored are private and only accessible by your team.</p>
            </div>
            <div>
              <h2>Command Line First</h2>
              <p className="hug">Index, search, and download documents from a CLI.</p>
            </div>
          </div>
        </div>
        <div className="showcase full-width alt">
          <div className="container cta-footer">
            <div>
              <h2>Try Starship For Free</h2>
            </div>
            <div className="actions">
              <a className="btn red">Get Started</a>
              <a className="btn green">Starship in 60 seconds</a>
            </div>
          </div>
        </div>
        <style jsx global>{`
          .pull-front, .pull-front * {
            z-index: 999;
          }

          .stars, .twinkling, .clouds {
            position:absolute;
            display:block;
            top:0; bottom:0;
            left:0; right:0;
            width:100%; height:100%;
          }

          .stars {
            z-index: 0;
            background: #FFF url('/static/stars.png') repeat top center;
          }

          .twinkling{
            z-index: 1;
            background:transparent url('/static/twinkling.png') repeat top center;
            animation: move-twink-back 200s linear infinite;
          }

          .clouds{
            z-index: 2;
            background:transparent url('/static/clouds.png') repeat top center;
            animation: move-clouds-back 200s linear infinite;
          }

          @keyframes move-twink-back {
            from {background-position:0 0;}
            to {background-position:-10000px 5000px;}
          }

          @keyframes move-clouds-back {
            from {background-position:0 0;}
            to {background-position:10000px 0;}
          }

          [data-terminal] {
              max-width: 100%;
              background: var(--color-bg);
              border-radius: 4px;
              padding: 40px 15px 20px 15px;
              position: relative;
              -webkit-box-sizing: border-box;
                      box-sizing: border-box;
          }

          [data-terminal]:before {
              content: '';
              position: absolute;
              top: 15px;
              left: 15px;
              display: inline-block;
              width: 15px;
              height: 15px;
              border-radius: 50%;
              /* A little hack to display the window buttons in one pseudo element. */
              background: #d9515d;
              -webkit-box-shadow: 25px 0 0 #f4c025, 50px 0 0 #3ec930;
                      box-shadow: 25px 0 0 #f4c025, 50px 0 0 #3ec930;
          }

          .cta-footer {
            display: grid;
            grid-template-columns: .75fr 1.25fr;
            justify-items: center;
            align-items: center;
          }

          .cta-footer h2 {
            text-align: left !important;
            margin: 0;
          }

          .cta-footer .btn:first-child {
            margin-right: 20px;
          }

          .cta-footer .btn {
            box-shadow: 0 20px 50px 0 rgba(0,0,0,0.2);
          }

          .hero .btn {
            box-shadow: 0 20px 50px 0 rgba(0,0,0,0.2);
          }

          .three {
            display: grid;
            grid-template-columns: 1fr 1fr 1fr;
            column-gap: 1.5rem;
          }

          .three div {
            justify-self: center;
          }

          .tab-section {
            display: grid;
            grid-template-columns: 1fr;
          }

          .three h2 {
            margin: 0;
            text-align: left !important;
            font-size: 100% !important;
          }

          .tab-section ul {
            justify-self: center;
            list-style: none;
            margin: 0;
            padding: 0;
          }

          .tab-section li {
            display: inline-block;
            float: left;
            padding: 5px 20px;
          }

          .tab-section li.active {
            border-bottom: 1px solid #122239;
          }

          .tacb-section li:hover {
            cursor: pointer;
          }

          .tab-section button {
            outline: none;
            border: none;
            color: #122239;
          }

          .tab-section button:hover {
            color: #122239;
            cursor: pointer;
          }

          .tab-section button {
            font-size: 90% !important;
          }

          .two-sided {
            margin: 3rem 0;
            display: grid;
            grid-template-columns: 1fr 1fr;
            column-gap: 1rem;
          }

          .two-sided .placeholder {
            height: 100%;
            width: 100%;
            background: #122239;
          }

          .two-sided p {
            margin: 1rem 0;
          }

          .hero {
            z-index: 999;
            padding: 100px 0 150px 0;
            width: 90%;
            position: relative;
            left: 5%;
            font-size: 132%;
            text-align: center;
          }

          .hero h2 {
            font-size: 2.121rem;
            margin-bottom: 1.5rem;
          }

          .hero-actions {
            font-size: 84%;
            margin: 40px 0 0 0;
            display: grid;
            column-gap: 1rem;
            grid-template-columns: 1fr 1fr;
            align-items: center;
          }

          .hero-actions div:first-child {
            justify-self: end;
          }

          .hero-actions div:last-child {
            justify-self: start;
          }

          .showcase {
            background: #fafbfc;
            padding: 3rem 0;
            border-bottom: 1px solid #eaeaea;
          }

          .showcase h2 {
            text-align: center;
          }

          .showcase.alt {
            background: #fff;
          }

          .showcase.alt p {
            color: #707d8e;
          }

          .full-width {
            width: 100vw;
            position: relative;
            left: 50%;
            right: 50%;
            margin-left: -50vw;
            margin-right: -50vw;
          }

          .long-width {
            position: relative;
            width: calc(100% + 146px);
            margin-left: -73px;
            margin-right: -73px;
          }

          .terminal {
            margin: 0 auto;
            background: #122239;
            box-shadow: 0 20px 50px 0 rgba(0,0,0,0.2);
            border-radius: 5px;
          }

          .terminal img {
            width: 100%;
          }
        `}</style>
      </>
    )
  }
}

// filter redux state to props for this page
const mapStateToProps = ({ posts }) => ({ posts })

// connect redux state to component
export default connect(
  mapStateToProps,
  null
)(Home)