var React = require('react');
var BootStrap = require('bootstrap');

var DataTypeSelection = require('./UI/DataTypeSelection.jsx');
var DisplayGraph = require('./UI/DisplayGraph.jsx');
var FacebookData = require('./UI/FacebookData.jsx');

var Main = React.createClass({
    getInitialState: function(){
        return{ active: [], login: false}
    },
    updateDisplayGraph: function(newDataTypes){
        this.setState({active: newDataTypes});
        console.log("Updated Data Types: " + this.state.active);
    },
    render: function(){
        return(
            <div>
                <div className="row">
                    <div className="col-md-3 col-md-offset-1">
                        <DataTypeSelection updateDisplayGraph={this.updateDisplayGraph}/>
                    </div>
                    <div className="col-md-6 col-md-offset-1">
                        <FacebookData loginStatus={this.state.login}/>
                    </div>
                </div>
                <div className="row">
                    <div className="col-md-8 col-md-offset-2">
                        <DisplayGraph dataTypes={this.state.active}/>
                    </div>
                </div>
            </div>
        );
    }
});

module.exports = Main;
