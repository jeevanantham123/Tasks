import { Route, Switch,Router} from 'react-router-dom';
import React from 'react';
import history from './History';
import Homepage from './components/Homepage';
import App from './App';

class Routes extends React.Component{
render(){
    return(
        <Router history={history}>
                <Switch>
                    <Route  path="/instagram-replica/" exact component={App} /> 
                    <Route path="/instagram-replica/home/"  component={Homepage} /> 
                </Switch>
        </Router>
    )
}
}
export default Routes;