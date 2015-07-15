var React = require('react');
var Bootstrap = require('bootstrap');
var $ = require('jquery');
var Content = require('./Content.jsx');

var FacebookData = React.createClass({
    getInitialState: function(){
        return{selected: '', contentData: ''}
    },
    componentDidMount: function(){
        $('#posts').prop('disabled', true);
        $('#pictures').prop('disabled', true);
        $('#videos').prop('disabled', true);
        $('#status').prop('disabled', true);
    },
    componentWillReceiveProps(nextProps){
        if(nextProps.loginStatus){
            $('#posts').prop('disabled', false);
            $('#pictures').prop('disabled', false);
            $('#videos').prop('disabled', false);
            $('#status').prop('disabled', false);
        }
    },
    renderData: function(data){
        this.setState({contentData: data});
    },
    loadPosts: function(event){
        this.setState({
             selected: 'posts'
        });
        $.get("/FBContent", {data:'posts'}, function(newdata){this.renderData(newdata);}.bind(this),"text");
    },
    loadPictures: function(event){
        this.setState({
            selected: 'pictures'
        });
        $.get("/FBContent", {data: 'pictures'}, function(newdata){this.renderData(newdata);}.bind(this), "text");
    },
    loadVideos: function(event){
        this.setState({
            selected: 'videos'
        });
        $.get("/FBContent", {data: 'videos'}, function(newdata){this.renderData(newdata);}.bind(this), "text");
    },
    loadStatus: function(event){
        this.setState({
            selected: 'status'
        });
        $.get("/FBContent", {data: 'status'}, function(newdata){this.renderData(newdata);}.bind(this), "text");
    },
    getButtonID: function(id){
        this.props.getButtonID(id);
    },
    render: function(){
        var errorLogin;
        if(!this.props.loginStatus){
            errorLogin = <h3 className="text-center"><span className="label label-danger">Please Login To Facebook</span></h3>
        }

        var loadContent;
        if(this.state.selected.length != 0){
            loadContent = <Content data={this.state.contentData} getButtonID={this.getButtonID}/>
        }

        return(
            <div className="panel panel-default">
                <div className="panel-body scroll-panel">
                    <div className="btn-group" role="group">
                        <button id="posts" name="posts" type="button" className="btn btn-success" onClick={this.loadPosts}>Posts</button>
                        <button id="pictures" name="pictures" type="button" className="btn btn-success" onClick={this.loadPictures}>Pictures</button>
                        <button id="videos" name="videos" type="button" className="btn btn-success" onClick={this.loadVideos}>Videos</button>
                        <button id="status" name="status" type="button" className="btn btn-success" onClick={this.loadStatus}>Status</button>
                    </div>
                    {errorLogin}
                    {loadContent}
                </div>
            </div>
        );
    }
});

module.exports = FacebookData;
