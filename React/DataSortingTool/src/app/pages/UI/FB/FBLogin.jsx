var React = require('react');
var Bootstrap = require('bootstrap');
var $ = require('jquery');

/*
 *  Class: FBLogin
 *  Usage: A Simple Button To Excute Login Aceess Into Facebook And Send The Response Object To The Server
 *  Design:
 *      -A Simple Button To Send A Post Request To The Server, where graph API calls will be done for the user
 *      -handleClick is the function used to start the FB login request and to post the information returned from facebook to the server
 */

var FBLogin = React.createClass({
    getInitialState: function(){
        return{login: true}
    },
    propTypes: {
        loginStatus: React.PropTypes.func.isRequired
    },
    handleClick: function(event){
        FB.login(function(response){
            console.log("Response: " + response);
            $.post('/FBToken', {object: response})
        }, {
            scope: 'user_videos,user_photos,user_posts,user_status',
            return_scopes: true,
            auth_type: 'rerequest'
        });

        this.setState({
            login: !this.state.login
        });

        console.log("FBLogin: " + this.state.login);
        this.props.loginStatus(this.state.login);
        $('#FBLogin').prop('disabled', true);
    },
    render: function(){
        return(
            <button id="FBLogin" name="login" type="button" className="btn btn-success" onClick={this.handleClick}>FaceBook Login</button>
        );
    },
});

module.exports = FBLogin;
