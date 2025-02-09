Module beamlit.models.agent_history_event
=========================================

Classes
-------

`AgentHistoryEvent(end: beamlit.types.Unset | str = <beamlit.types.Unset object>, error: beamlit.types.Unset | str = <beamlit.types.Unset object>, name: beamlit.types.Unset | str = <beamlit.types.Unset object>, parameters: beamlit.types.Unset | str = <beamlit.types.Unset object>, start: beamlit.types.Unset | str = <beamlit.types.Unset object>, status: beamlit.types.Unset | str = <beamlit.types.Unset object>, sub_function: beamlit.types.Unset | str = <beamlit.types.Unset object>, took: beamlit.types.Unset | int = <beamlit.types.Unset object>, type_: beamlit.types.Unset | str = <beamlit.types.Unset object>)`
:   Agent deployment history event
    
    Attributes:
        end (Union[Unset, str]): End time
        error (Union[Unset, str]): Error message
        name (Union[Unset, str]): Name of the function or agent
        parameters (Union[Unset, str]): Parameters
        start (Union[Unset, str]): Start time
        status (Union[Unset, str]): Status, eg: running, success, failed
        sub_function (Union[Unset, str]): Function used in kit if a kit was used
        took (Union[Unset, int]): Number of milliseconds it took to complete the event
        type_ (Union[Unset, str]): Type, one of function or agent
    
    Method generated by attrs for class AgentHistoryEvent.

    ### Static methods

    `from_dict(src_dict: dict[str, typing.Any]) ‑> ~T`
    :

    ### Instance variables

    `additional_keys: list[str]`
    :

    `additional_properties: dict[str, typing.Any]`
    :

    `end: beamlit.types.Unset | str`
    :

    `error: beamlit.types.Unset | str`
    :

    `name: beamlit.types.Unset | str`
    :

    `parameters: beamlit.types.Unset | str`
    :

    `start: beamlit.types.Unset | str`
    :

    `status: beamlit.types.Unset | str`
    :

    `sub_function: beamlit.types.Unset | str`
    :

    `took: beamlit.types.Unset | int`
    :

    `type_: beamlit.types.Unset | str`
    :

    ### Methods

    `to_dict(self) ‑> dict[str, typing.Any]`
    :