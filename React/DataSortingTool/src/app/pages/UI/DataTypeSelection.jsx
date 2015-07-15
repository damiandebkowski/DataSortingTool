var React = require('react');
var Bootstrap = require('bootstrap');
var CheckBox = require('./CheckBox.jsx');
var $ = require('jquery');

/*
 * Class: DataTypeSelection
 * Usage: Allow the user to click checkBoxes to select the data types that they wish to analyze in the DisplayGraph Class
 * Design:
 *      -Var Array: active[] to store checkbox ids created in this UI
 *      -function updateDataType
 *          Parameter var: DataType is the propID from the checkbox that should be added to or removed from array active based on the value of method
 *          if method is true the checkbox is checked then add the propID
 *          if false the checkbox is unchecked then remove the propID
 *          edge case if the array size is 1, then remove the last element
 *          then calls updateDisplayGraph
 *      -function callback updateDisplayGraph is required
 *          -sends array active back to parent node, so that the DisplayGraph can be updated
 */

var DataTypeSelection = React.createClass({
    getInitialState: function(){
        return{active: []}
    },
    propTypes: {
        updateDisplayGraph: React.PropTypes.func.isRequired
    },
    updateDataType: function(DataType, method){
        var temp = this.state.active
        if(method){
            temp.push(DataType)
        }else if(this.state.active.length != 1){
            temp.splice(temp.indexOf(DataType), 1)
        }else{
            temp.pop()
        }

        this.setState({active: temp});
        temp = temp.toString();
        this.props.updateDisplayGraph(this.state.active);
    },
    render: function(){
        return(
            <div className="panel panel-default">
                <div className="panel-body">
                    <form className="form-horizontal" action="/SortData" method="post">
                        <div className="form-group">
                            <div className="col-sm-12">
                                <h3 className="text-center">
                                    <span className="label label-success">Select Data Types:</span>
                                </h3>
                            </div>
                        </div>
                        <div className="form-group">
                            <div className="col-sm-1 col-sm-offset-2">
                                <CheckBox id="Gender" name="Gender" checkDataType={this.updateDataType}/>
                            </div>
                            <div className="col-sm-5">
                                <label>Gender</label>
                            </div>
                        </div>
                        <div className="form-group">
                            <div className="col-sm-1 col-sm-offset-2">
                                <CheckBox id="Age" name="Age" checkDataType={this.updateDataType}/>
                            </div>
                            <div className="col-sm-5">
                                <label>Age</label>
                            </div>
                        </div>
                        <div className="form-group">
                            <div className="col-sm-1 col-sm-offset-2">
                                <CheckBox id="Location" name="Location" checkDataType={this.updateDataType}/>
                            </div>
                            <div className="col-sm-5">
                                <label>Location</label>
                            </div>
                        </div>
                        <div className="form-group">
                            <div className="col-sm-1 col-sm-offset-2">
                                <CheckBox id="Education" name="Education" checkDataType={this.updateDataType}/>
                            </div>
                            <div className="col-sm-5">
                                <label>Education</label>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        );
    }
});

module.exports = DataTypeSelection;
