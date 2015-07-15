var React = require('react');
var Bootstrap = require('bootstrap');
var $ = require('jquery');

var Content = React.createClass({
    render: function(){
        var items = [];

        for(var i = 0; i < 20; i++){
            items.push(<button type="button" className="list-group-item">Button {i}</button>)
        }

        return(
            <div className="list-group">
                {items}
            </div>
        );
    }
});

module.exports = Content;
