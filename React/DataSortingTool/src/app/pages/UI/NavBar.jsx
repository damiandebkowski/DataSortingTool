var React = require('react');
var Bootstrap = require('bootstrap');
var FBLoginButton = require('./FB/FBLogin.jsx');

/*
 * Class: NavBar
 * Usage: Navigate between the pages of the app and contains the facebook login button to obtain access_tokens
 * Design:
 *      -Provide navigation back to the homepage = main app function
 *      -To the About page for information on how to use the app and what it is designed for
 *      -To the Contact page for information to get into contact with me incase of issues
 *      -To the private policy page for legal reasons
*/

var NavBar = React.createClass({
    getLoginStatus: function(status){
        this.props.updateLoginStatus(status);
    },
    propTypes: {
        updateLoginStatus: React.PropTypes.func.isRequired,
        getRenderState: React.PropTypes.func.isRequired
    },
    updateRenderState: function(renderState){
        window.location.href='#' + renderState;
        this.props.getRenderState(renderState);
    },
    render: function(){
        return(
            <nav className="navbar navbar-default">
                <div className="container">
                    <div className="navbar-header">
                        <a className="navbar-brand" href="#">Data Sorting</a>
                    </div>
                    <div id="navbar" className="collapse navbar-collapse">
                        <button type="button" className="navbar-btn btn btn-default" onClick={this.updateRenderState.bind(this, 'HOME')}>Home</button>
                        <button type="button" className="navbar-btn btn btn-default" onClick={this.updateRenderState.bind(this, 'ABOUT')}>About</button>
                        <button type="button" className="navbar-btn btn btn-default" onClick={this.updateRenderState.bind(this, 'CONTACT')}>Contact</button>
                        <button href="#policy" type="button" className="navbar-btn btn btn-default" onClick={this.updateRenderState.bind(this, 'POLICY')}>Private Policy</button>
                        <form id="facebookLogin" name="facebookLogin" className="navbar-form navbar-right" action="/FBToken" method="post">
                            <FBLoginButton loginStatus={this.getLoginStatus}/>
                        </form>
                    </div>
                </div>
            </nav>
        );
    }
});

module.exports = NavBar;
