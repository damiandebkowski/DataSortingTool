var React = require('react');
var Bootstrap = require('bootstrap');
var $ = require('jquery');
var Button = require('./ContentButton.jsx');

var Content = React.createClass({
    shouldComponentUpdate: function(nextProps, nextState){
        return this.props.data != nextProps.data;
    },
    getButtonID: function(id){
        this.props.getButtonID(id);
    },
    render: function(){
        console.log("Rendering Contents From Facebook " + this.props.data)

        var items = [];

        var serverData = this.props.data;
        var itemIDs = serverData.split("/");
        for(var i = 0; i < itemIDs.length; i++){
            items.push(<Button id={itemIDs[i]} getButtonID={this.getButtonID}/>)
        }

        return(
            <div className="list-group">
                {items}
            </div>
        );
    }
});

module.exports = Content;
