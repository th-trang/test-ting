import Head from "next/head"
import Page from "./page"

import Providers from './providers'

export const metadata = {
  title: 'KOL Dashboard',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <>
      <Head>
        <title>{metadata.title}</title>
      </Head>
      <div className="dark">
        <Providers>
          {children}
        </Providers>
      </div>
    </>
  );
}
