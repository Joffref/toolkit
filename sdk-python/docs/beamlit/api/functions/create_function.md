Module beamlit.api.functions.create_function
============================================

Functions
---------

`asyncio(*, client: beamlit.client.AuthenticatedClient, body: beamlit.models.function.Function) ‑> beamlit.models.function.Function | None`
:   Create function
    
    Args:
        body (Function): Function
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Function

`asyncio_detailed(*, client: beamlit.client.AuthenticatedClient, body: beamlit.models.function.Function) ‑> beamlit.types.Response[beamlit.models.function.Function]`
:   Create function
    
    Args:
        body (Function): Function
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Function]

`sync(*, client: beamlit.client.AuthenticatedClient, body: beamlit.models.function.Function) ‑> beamlit.models.function.Function | None`
:   Create function
    
    Args:
        body (Function): Function
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Function

`sync_detailed(*, client: beamlit.client.AuthenticatedClient, body: beamlit.models.function.Function) ‑> beamlit.types.Response[beamlit.models.function.Function]`
:   Create function
    
    Args:
        body (Function): Function
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Function]