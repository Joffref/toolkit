Module beamlit.api.models.get_model
===================================

Functions
---------

`asyncio(model_name: str, *, client: beamlit.client.AuthenticatedClient, environment: beamlit.types.Unset | str = <beamlit.types.Unset object>) ‑> beamlit.models.model.Model | None`
:   Get model
    
     Returns a model by name.
    
    Args:
        model_name (str):
        environment (Union[Unset, str]):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Model

`asyncio_detailed(model_name: str, *, client: beamlit.client.AuthenticatedClient, environment: beamlit.types.Unset | str = <beamlit.types.Unset object>) ‑> beamlit.types.Response[beamlit.models.model.Model]`
:   Get model
    
     Returns a model by name.
    
    Args:
        model_name (str):
        environment (Union[Unset, str]):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Model]

`sync(model_name: str, *, client: beamlit.client.AuthenticatedClient, environment: beamlit.types.Unset | str = <beamlit.types.Unset object>) ‑> beamlit.models.model.Model | None`
:   Get model
    
     Returns a model by name.
    
    Args:
        model_name (str):
        environment (Union[Unset, str]):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Model

`sync_detailed(model_name: str, *, client: beamlit.client.AuthenticatedClient, environment: beamlit.types.Unset | str = <beamlit.types.Unset object>) ‑> beamlit.types.Response[beamlit.models.model.Model]`
:   Get model
    
     Returns a model by name.
    
    Args:
        model_name (str):
        environment (Union[Unset, str]):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Model]