var React = require('react');
var Bootstrap = require('bootstrap');
var NavBar = require('./pages/UI/NavBar.jsx');
var Main = require('./pages/main.jsx');
var About = require('./pages/About.jsx');
var Contact = require('./pages/Contact.jsx');
var Policy = require('./pages/Policy.jsx');

/*
 * Class: Container
 * Usage: This is the main UI container, that holds all of the other sub UI containers, such as DisplayGraph and DataTypeSelection
 * Design:
 *      -Sub Containers added here must contain function callback in order to return updated information from child nodes to be passed back up to this container, so that other child nodes from here
 *          can utilize information from other seperate sub-containers (ex. between DataTypeSelection and DisplayGraph)
 */

var Container = React.createClass({
    getInitialState: function(){
        return{login: false, renderState: 'HOME'}
    },
    updateLoginStatus: function(status){
        this.setState({login: status});
    },
    getRenderState: function(newRenderState){
        console.log("Container SetRenderState: " + newRenderState);
        this.setState({renderState: newRenderState});
    },
    render: function(){
        var render;
        switch(this.state.renderState){
            case 'HOME':
                console.log("Rendering Home")
                render = <Main loginStatus={this.state.login}/>
                break;
            case 'ABOUT':
                console.log("Rendering About");
                render = <About/>
                break;
            case 'CONTACT':
                console.log("Rendering Contact");
                render = <Contact/>
                break;
            case 'POLICY':
                console.log("Rendering Private Policy");
                render = <Policy/>
            break;
            default:
                console.log("Error: renderState invalid");
        }
        return(
            <div>
                <div className="row">
                    <div className="col-md-12">
                        <NavBar getRenderState={this.getRenderState} updateLoginStatus={this.updateLoginStatus}/>
                    </div>
                </div>
                {render}
            </div>
        );
    }
});

module.exports = Container;
