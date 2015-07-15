var React = require('react');
var Bootstrap = require('bootstrap');
var $ = require('jquery');

var ContentButton = React.createClass({
    handleClick: function(event){
        this.props.getButtonID(this.props.id)
    },
    render: function(){
        return(
            <button type="button" key={this.props.id} className="list-group-item" onClick={this.handleClick}>Button ID: {this.props.id}</button>
        );
    }
});

module.exports = ContentButton;
