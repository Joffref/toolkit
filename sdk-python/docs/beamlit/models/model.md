Module beamlit.models.model
===========================

Classes
-------

`Model(events: beamlit.types.Unset | list['CoreEvent'] = <beamlit.types.Unset object>, metadata: beamlit.types.Unset | ForwardRef('EnvironmentMetadata') = <beamlit.types.Unset object>, spec: beamlit.types.Unset | ForwardRef('ModelSpec') = <beamlit.types.Unset object>, status: beamlit.types.Unset | str = <beamlit.types.Unset object>)`
:   Logical object representing a model, that can be instantiated in multiple environments as model deployments
    
    Attributes:
        events (Union[Unset, list['CoreEvent']]): Core events
        metadata (Union[Unset, EnvironmentMetadata]): Environment metadata
        spec (Union[Unset, ModelSpec]): Model specification
        status (Union[Unset, str]): Model status
    
    Method generated by attrs for class Model.

    ### Static methods

    `from_dict(src_dict: dict[str, typing.Any]) ‑> ~T`
    :

    ### Instance variables

    `additional_keys: list[str]`
    :

    `additional_properties`
    :

    `events`
    :

    `metadata`
    :

    `spec`
    :

    `status`
    :

    ### Methods

    `to_dict(self) ‑> dict[str, typing.Any]`
    :