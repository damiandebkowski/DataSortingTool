var React = require('react');
var BootStrap = require('bootstrap');

var DataTypeSelection = require('./UI/DataTypeSelection.jsx');
var DisplayGraph = require('./UI/DisplayGraph.jsx');
var FacebookData = require('./UI/FacebookData.jsx');

var Main = React.createClass({
    getInitialState: function(){
        return{ active: [], buttonID: ''}
    },
    updateDisplayGraph: function(newDataTypes){
        this.setState({active: newDataTypes});
    },
    getButtonID: function(id){
        console.log("Main Button ID: " + id)
        this.setState({
            buttonID: id
        });
    },
    render: function(){
        return(
            <div>
                <div className="row">
                    <div className="col-md-3 col-md-offset-1">
                        <DataTypeSelection updateDisplayGraph={this.updateDisplayGraph}/>
                    </div>
                    <div className="col-md-6 col-md-offset-1">
                        <FacebookData loginStatus={this.props.loginStatus} getButtonID={this.getButtonID}/>
                    </div>
                </div>
                <div className="row">
                    <div className="col-md-8 col-md-offset-2">
                        <DisplayGraph dataTypes={this.state.active} buttonID={this.state.buttonID}/>
                    </div>
                </div>
            </div>
        );
    }
});

module.exports = Main;
