import Document, { Head, Main, NextScript } from 'next/document';
import flush from 'styled-jsx/server'
import DocumentHead from '../components/shared/documenthead'

export default class MyDocument extends Document {
  static async getInitialProps ({ renderPage }) {
    const { html, head } = renderPage()
    const styles = flush()
    return { html, head, styles }
  }

  render() {
    return (
      <html lang="en">
        <Head>
          <DocumentHead />
        </Head>
        <body>
          <Main />
          <NextScript />
        </body>
      </html>
    );
  }
}