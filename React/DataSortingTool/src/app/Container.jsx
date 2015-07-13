var React = require('react');
var Bootstrap = require('bootstrap');
var NavBar = require('./UI/NavBar.jsx');
var DataTypeSelection = require('./UI/DataTypeSelection.jsx');
var DisplayGraph = require('./UI/DisplayGraph.jsx');
var FacebookData = require('./UI/FacebookData.jsx');

/*
 * Class: Container
 * Usage: This is the main UI container, that holds all of the other sub UI containers, such as DisplayGraph and DataTypeSelection
 * Design:
 *      -Sub Containers added here must contain function callback in order to return updated information from child nodes to be passed back up to this container, so that other child nodes from here
 *          can utilize information from other seperate sub-containers (ex. between DataTypeSelection and DisplayGraph)
 */

var Container = React.createClass({
    getInitialState: function(){
        return{active: [], login: false, renderState: 'HOME'}
    },
    //This is a callback function passed down to DataTypeSelection To Retrieve data back in newDataTypes
    //newDataTypes is a copy of the var active in DataTypeSelection and will be used to update this.active and pass down this.active to DisplayGraph as a prop
    updateDisplayGraph: function(newDataTypes){
        this.setState({
            active: newDataTypes
        });
        console.log("Display Graph Updated");
    },
    updateLoginStatus: function(status){
        this.setState({
            login: status
        });
    },
    getRenderState: function(newRenderState){
        console.log("Container SetRenderState: " + newRenderState);
        this.setState({renderState: newRenderState});
    },
    render: function(){
        var render;
        switch(this.state.renderState){
            case 'HOME':
                break;
            case 'ABOUT':
                break;
            case 'CONTACT':
                break;
            case 'POLICY':
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

module.exports = Container;
