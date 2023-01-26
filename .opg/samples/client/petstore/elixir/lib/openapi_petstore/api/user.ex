# NOTE: This file is auto generated by OpenAPI Generator 6.3.0-SNAPSHOT (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule OpenapiPetstore.Api.User do
  @moduledoc """
  API calls for all endpoints tagged `User`.
  """

  alias OpenapiPetstore.Connection
  import OpenapiPetstore.RequestBuilder

  @doc """
  Create user
  This can only be done by the logged in user.

  ### Parameters

  - `connection` (OpenapiPetstore.Connection): Connection to server
  - `user` (User): Created user object
  - `opts` (keyword): Optional parameters

  ### Returns

  - `{:ok, nil}` on success
  - `{:error, Tesla.Env.t}` on failure
  """
  @spec create_user(Tesla.Env.client, OpenapiPetstore.Model.User.t, keyword()) :: {:ok, nil} | {:error, Tesla.Env.t}
  def create_user(connection, user, _opts \\ []) do
    request =
      %{}
      |> method(:post)
      |> url("/user")
      |> add_param(:body, :body, user)
      |> Enum.into([])

    connection
    |> Connection.request(request)
    |> evaluate_response([
      {:default, false}
    ])
  end

  @doc """
  Creates list of users with given input array
  

  ### Parameters

  - `connection` (OpenapiPetstore.Connection): Connection to server
  - `user` ([OpenapiPetstore.Model.User.t]): List of user object
  - `opts` (keyword): Optional parameters

  ### Returns

  - `{:ok, nil}` on success
  - `{:error, Tesla.Env.t}` on failure
  """
  @spec create_users_with_array_input(Tesla.Env.client, list(OpenapiPetstore.Model.User.t), keyword()) :: {:ok, nil} | {:error, Tesla.Env.t}
  def create_users_with_array_input(connection, user, _opts \\ []) do
    request =
      %{}
      |> method(:post)
      |> url("/user/createWithArray")
      |> add_param(:body, :body, user)
      |> Enum.into([])

    connection
    |> Connection.request(request)
    |> evaluate_response([
      {:default, false}
    ])
  end

  @doc """
  Creates list of users with given input array
  

  ### Parameters

  - `connection` (OpenapiPetstore.Connection): Connection to server
  - `user` ([OpenapiPetstore.Model.User.t]): List of user object
  - `opts` (keyword): Optional parameters

  ### Returns

  - `{:ok, nil}` on success
  - `{:error, Tesla.Env.t}` on failure
  """
  @spec create_users_with_list_input(Tesla.Env.client, list(OpenapiPetstore.Model.User.t), keyword()) :: {:ok, nil} | {:error, Tesla.Env.t}
  def create_users_with_list_input(connection, user, _opts \\ []) do
    request =
      %{}
      |> method(:post)
      |> url("/user/createWithList")
      |> add_param(:body, :body, user)
      |> Enum.into([])

    connection
    |> Connection.request(request)
    |> evaluate_response([
      {:default, false}
    ])
  end

  @doc """
  Delete user
  This can only be done by the logged in user.

  ### Parameters

  - `connection` (OpenapiPetstore.Connection): Connection to server
  - `username` (String.t): The name that needs to be deleted
  - `opts` (keyword): Optional parameters

  ### Returns

  - `{:ok, nil}` on success
  - `{:error, Tesla.Env.t}` on failure
  """
  @spec delete_user(Tesla.Env.client, String.t, keyword()) :: {:ok, nil} | {:error, Tesla.Env.t}
  def delete_user(connection, username, _opts \\ []) do
    request =
      %{}
      |> method(:delete)
      |> url("/user/#{username}")
      |> Enum.into([])

    connection
    |> Connection.request(request)
    |> evaluate_response([
      {400, false},
      {404, false}
    ])
  end

  @doc """
  Get user by user name
  

  ### Parameters

  - `connection` (OpenapiPetstore.Connection): Connection to server
  - `username` (String.t): The name that needs to be fetched. Use user1 for testing.
  - `opts` (keyword): Optional parameters

  ### Returns

  - `{:ok, OpenapiPetstore.Model.User.t}` on success
  - `{:error, Tesla.Env.t}` on failure
  """
  @spec get_user_by_name(Tesla.Env.client, String.t, keyword()) :: {:ok, nil} | {:ok, OpenapiPetstore.Model.User.t} | {:error, Tesla.Env.t}
  def get_user_by_name(connection, username, _opts \\ []) do
    request =
      %{}
      |> method(:get)
      |> url("/user/#{username}")
      |> Enum.into([])

    connection
    |> Connection.request(request)
    |> evaluate_response([
      {200, %OpenapiPetstore.Model.User{}},
      {400, false},
      {404, false}
    ])
  end

  @doc """
  Logs user into the system
  

  ### Parameters

  - `connection` (OpenapiPetstore.Connection): Connection to server
  - `username` (String.t): The user name for login
  - `password` (String.t): The password for login in clear text
  - `opts` (keyword): Optional parameters

  ### Returns

  - `{:ok, String.t}` on success
  - `{:error, Tesla.Env.t}` on failure
  """
  @spec login_user(Tesla.Env.client, String.t, String.t, keyword()) :: {:ok, nil} | {:ok, String.t} | {:error, Tesla.Env.t}
  def login_user(connection, username, password, _opts \\ []) do
    request =
      %{}
      |> method(:get)
      |> url("/user/login")
      |> add_param(:query, :username, username)
      |> add_param(:query, :password, password)
      |> Enum.into([])

    connection
    |> Connection.request(request)
    |> evaluate_response([
      {200, false},
      {400, false}
    ])
  end

  @doc """
  Logs out current logged in user session
  

  ### Parameters

  - `connection` (OpenapiPetstore.Connection): Connection to server
  - `opts` (keyword): Optional parameters

  ### Returns

  - `{:ok, nil}` on success
  - `{:error, Tesla.Env.t}` on failure
  """
  @spec logout_user(Tesla.Env.client, keyword()) :: {:ok, nil} | {:error, Tesla.Env.t}
  def logout_user(connection, _opts \\ []) do
    request =
      %{}
      |> method(:get)
      |> url("/user/logout")
      |> Enum.into([])

    connection
    |> Connection.request(request)
    |> evaluate_response([
      {:default, false}
    ])
  end

  @doc """
  Updated user
  This can only be done by the logged in user.

  ### Parameters

  - `connection` (OpenapiPetstore.Connection): Connection to server
  - `username` (String.t): name that need to be deleted
  - `user` (User): Updated user object
  - `opts` (keyword): Optional parameters

  ### Returns

  - `{:ok, nil}` on success
  - `{:error, Tesla.Env.t}` on failure
  """
  @spec update_user(Tesla.Env.client, String.t, OpenapiPetstore.Model.User.t, keyword()) :: {:ok, nil} | {:error, Tesla.Env.t}
  def update_user(connection, username, user, _opts \\ []) do
    request =
      %{}
      |> method(:put)
      |> url("/user/#{username}")
      |> add_param(:body, :body, user)
      |> Enum.into([])

    connection
    |> Connection.request(request)
    |> evaluate_response([
      {400, false},
      {404, false}
    ])
  end
end