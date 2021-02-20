import {useState} from 'react';
import Navbar from "./components/Navbar/Navbar";
import Sidebar from "./components/Sidebar/Sidebar";
import Main from './components/Main/Main';
import { ApolloClient, InMemoryCache, ApolloProvider } from '@apollo/client';

const client = new ApolloClient({
  uri: process.env.GQL_ENDPOINT,
  cache: new InMemoryCache()
});

const App = () => {

  const [sidebarOpen, setSidebarOpen] = useState(false);
  const openSidebar = () => {
    setSidebarOpen(true);
  }

  const closeSidebar = () => {
    setSidebarOpen(false);
  }

  return(
    <ApolloProvider client={client}>
      <div className="container">
        <Navbar sidebarOpen={sidebarOpen} openSidebar={openSidebar}/>
          <Main />
        <Sidebar sidebarOpen={sidebarOpen} closeSidebar={closeSidebar}/>
      </div>
    </ApolloProvider>
  );
}

export default App;
